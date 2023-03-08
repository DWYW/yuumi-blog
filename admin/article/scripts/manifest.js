const path = require('path')
const fs = require('fs')
const dotenv = require('dotenv')
const { program } = require('commander')
const { name, version } = require(path.resolve('package.json'))
const microName = path.parse(name).name

const options = program
  .option('--mode <mode>')
  .option('--version <version>')
  .parse(process.argv)
  .opts();

const envModeConfig = dotenv.parse(fs.readFileSync(`.env.${options.mode}`))

const manifest = {
  name: microName,
  latest: version,
  publicPath: envModeConfig.VUE_APP_PUBLIC_PATH,
  index: `${envModeConfig.VUE_APP_PUBLIC_PATH}${version}/index.html`
}

const buffer = Buffer.from(JSON.stringify(manifest, null, 2))
fs.writeFileSync(`dist/manifest.json`, buffer)
