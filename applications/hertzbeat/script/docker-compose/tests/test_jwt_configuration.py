# Licensed to the Apache Software Foundation (ASF) under one or more
# contributor license agreements.  See the NOTICE file distributed with
# this work for additional information regarding copyright ownership.
# The ASF licenses this file to You under the Apache License, Version 2.0
# (the "License"); you may not use this file except in compliance with
# the License.  You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

"""Check JWT configuration wiring in the bundled deployment files."""

import json
import os
from pathlib import Path
import shlex
import shutil
import subprocess
import tempfile
import unittest

import yaml


COMPOSE_ROOT = Path(__file__).resolve().parents[1]
DEPLOYMENTS = (
    "hertzbeat-mysql-iotdb",
    "hertzbeat-mysql-tdengine",
    "hertzbeat-mysql-victoria-metrics",
    "hertzbeat-postgresql-iotdb",
)


def application_config(path):
    return next(yaml.safe_load_all(path.read_text(encoding="utf-8")))


class JwtConfigurationTest(unittest.TestCase):
    def test_packaged_application_uses_generated_secret(self):
        config = application_config(COMPOSE_ROOT.parent / "application.yml")
        self.assertEqual(config["sureness"]["jwt"]["secret"], "")

    def test_compose_application_requires_environment_secret(self):
        for name in DEPLOYMENTS:
            with self.subTest(deployment=name):
                config = application_config(COMPOSE_ROOT / name / "conf/application.yml")
                self.assertEqual(config["sureness"]["jwt"]["secret"], "${SURENESS_JWT_SECRET}")

    def test_compose_passes_required_secret_to_manager(self):
        for name in DEPLOYMENTS:
            with self.subTest(deployment=name):
                compose = yaml.safe_load((COMPOSE_ROOT / name / "docker-compose.yaml").read_text())
                manager = compose["services"]["hertzbeat"]
                self.assertRegex(
                    manager["environment"].get("SURENESS_JWT_SECRET", ""),
                    r"^\$\{SURENESS_JWT_SECRET:\?[^{}]+\}$",
                )
                self.assertIn("./conf/application.yml:/opt/hertzbeat/config/application.yml", manager["volumes"])


@unittest.skipUnless(os.environ.get("COMPOSE") or shutil.which("docker"),
                     "Set COMPOSE or install Docker Compose to run config rendering tests")
class ComposeRenderingTest(unittest.TestCase):
    def render(self, deployment, secret=None):
        command = shlex.split(os.environ.get("COMPOSE", "docker compose"))
        command.extend([
            "--project-name", "jwt-test",
            "--env-file", os.devnull,
            "--file", str(COMPOSE_ROOT / deployment / "docker-compose.yaml"),
            "config", "--format", "json",
        ])
        with tempfile.TemporaryDirectory() as directory:
            environment = os.environ.copy()
            environment["DOCKER_CONFIG"] = directory
            environment.pop("SURENESS_JWT_SECRET", None)
            if secret is not None:
                environment["SURENESS_JWT_SECRET"] = secret
            return subprocess.run(command, capture_output=True, text=True, check=False,
                                  env=environment, timeout=30)

    def test_unset_secret_is_rejected(self):
        for name in DEPLOYMENTS:
            with self.subTest(deployment=name):
                result = self.render(name)
                self.assertNotEqual(result.returncode, 0)
                self.assertIn("SURENESS_JWT_SECRET", result.stderr)

    def test_empty_secret_is_rejected(self):
        for name in DEPLOYMENTS:
            with self.subTest(deployment=name):
                result = self.render(name, "")
                self.assertNotEqual(result.returncode, 0)
                self.assertIn("SURENESS_JWT_SECRET", result.stderr)

    def test_custom_secret_is_injected_unchanged(self):
        secret = "test-only: key # with \"quotes\" and literal\\slashes" * 2
        for name in DEPLOYMENTS:
            with self.subTest(deployment=name):
                result = self.render(name, secret)
                self.assertEqual(result.returncode, 0, result.stderr)
                config = json.loads(result.stdout)
                actual = config["services"]["hertzbeat"]["environment"]["SURENESS_JWT_SECRET"]
                self.assertEqual(actual, secret)


if __name__ == "__main__":
    unittest.main()
