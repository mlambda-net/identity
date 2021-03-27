const path = require('path');

module.exports = (config, env) => {



  config.module.rules.push(
    {
      test: /\.jsx$/,
      use: {
        loader: 'babel-loader'
      }
    })

  config.module.rules.push({
    test: /\.html$/,
    use: [
      {
        loader: 'html-loader',
      },
    ],
  })

  config.module.rules.push({
    test: /\.styl$/,
    use: [
      "style-loader",
      {
        loader: "stylus-loader",
      }
    ],
  })

  return config
}
