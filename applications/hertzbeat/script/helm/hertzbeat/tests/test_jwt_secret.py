# Licensed to the Apache Software Foundation (ASF) under one
# or more contributor license agreements.  See the NOTICE file
# distributed with this work for additional information
# regarding copyright ownership.  The ASF licenses this file
# to you under the Apache License, Version 2.0 (the
# "License"); you may not use this file except in compliance
# with the License.  You may obtain a copy of the License at
#
#   http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing,
# software distributed under the License is distributed on an
# "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
# KIND, either express or implied.  See the License for the
# specific language governing permissions and limitations
# under the License.

"""JWT configuration regression tests. Requires Helm 3+ and PyYAML."""

import json
import os
from pathlib import Path
import subprocess
import tempfile
import unittest

import yaml


CHART = Path(__file__).resolve().parents[1]
HELM = os.environ.get("HELM", "helm")
# Public, compromised defaults retained only as regression-test inputs.
LEGACY_SECRETS = (
    (
        'CyaFv0bwq2Eik0jdrKUtsA6bx3sDJeFV643R '
        'LnfKefTjsIfJLBa2YkhEqEGtcHDTNe4CU6+9 '
        '8tVt4bisXQ13rbN0oxhUZR73M6EByXIO+SV5 '
        'dKhaX0csgOCTlCxq20yhmUea6H6JIpSE2Rwp'
    ),
    (
        'CyaFv0bwq2Eik0jdrKUtsA6bx3sDJeFV143R '
        'LnfKefTjsIfJLBa2YkhEqEGtcHDTNe4CU6+9 '
        '8tVt4bisXQ13rbN0oxhUZR73M6EByXIO+SV5 '
        'dKhaX0csgOCTlCxq20yhmUea6H6JIpSE2Rwp'
    ),
)


class JwtSecretTest(unittest.TestCase):
    def render(self, values=None):
        command = [HELM, "template", "jwt-test", str(CHART), "--show-only",
                   "templates/manager/configmap.yaml"]
        with tempfile.TemporaryDirectory() as directory:
            if values is not None:
                values_path = Path(directory) / "values.json"
                values_path.write_text(json.dumps(values), encoding="utf-8")
                command.extend(["--values", str(values_path)])
            return subprocess.run(command, capture_output=True, text=True, check=False)

    def test_defaults_require_an_explicit_secret(self):
        result = self.render()
        self.assertNotEqual(result.returncode, 0)
        self.assertIn("manager.jwtSecretKey", result.stderr)

    def test_empty_secrets_are_rejected(self):
        for secret in (None, "", " ", " \t\n "):
            with self.subTest(secret=repr(secret)):
                result = self.render({"manager": {"jwtSecretKey": secret}})
                self.assertNotEqual(result.returncode, 0)
                self.assertIn("manager.jwtSecretKey", result.stderr)

    def test_both_public_legacy_secrets_are_rejected(self):
        for index, secret in enumerate(LEGACY_SECRETS):
            with self.subTest(legacy=index):
                result = self.render({"manager": {"jwtSecretKey": secret}})
                self.assertNotEqual(result.returncode, 0)
                self.assertIn("public default", result.stderr)

    def test_custom_secret_survives_both_yaml_layers(self):
        for secret in (
            "test-only-custom-secret-" * 4,
            "  test-only: value # with 'quotes' \"double\" \\ slash and\nnewline  ",
            "012345678901234567890123456789012345678901234567890123456789012345",
        ):
            with self.subTest(case=secret[:10]):
                result = self.render({"manager": {"jwtSecretKey": secret}})
                self.assertEqual(result.returncode, 0, result.stderr)
                configmap = yaml.safe_load(result.stdout)
                config = next(yaml.safe_load_all(configmap["data"]["application.yml"]))
                self.assertEqual(config["sureness"]["jwt"]["secret"], secret)


if __name__ == "__main__":
    unittest.main()
