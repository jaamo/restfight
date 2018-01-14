const path = require('path');

module.exports = {
  entry: {
    app: ['whatwg-fetch', './babel/index.js']
  },
  module: {
    loaders: [
      {
        test: /\.js$/,
        exclude: /node_modules/,
        loader: 'babel-loader',
        query: {
          presets: ['es2015']
        }
      }
    ]
  },
  output: {
    filename: 'index.js',
    path: path.resolve(__dirname, 'js')
  }
};