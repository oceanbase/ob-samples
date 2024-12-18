export PRISMA_ENGINES_MIRROR=https://oceanbase-prisma-builds.s3.ap-southeast-1.amazonaws.com
export BINARY_DOWNLOAD_VERSION=96fa66f2f130d66795d9f79dd431c678a9c7104e
npm install
npx prisma migrate dev --name init
npx ts-node index.ts
