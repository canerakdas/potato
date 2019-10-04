
var HtmlWebpackPlugin = require('html-webpack-plugin');
var Webpack = require('webpack');
var path = require('path');

module.exports = {
  entry: {
    main: path.join(__dirname, '{{.Source}}', '{{.DefaultJS}}')
  },
  output: {
    path: path.join(__dirname, 'dist'),
    filename: "[name]-bundle.js"
  },
  module: {
    rules: [
      {
        test: /\.(js|jsx)$/,
        exclude: /node_modules/,
        use: {
          loader: "babel-loader"
        }
      },
      {
        test: /\.css$/i,
        use: ['style-loader', 'css-loader'],
      },
    ]
  },
  plugins: [new HtmlWebpackPlugin(
    {
      chunks: ['main'],
      filename: '{{.DefaultHtml}}',
      title: "{{.Name}}",
      template: path.join(__dirname, '{{.Source}}', '{{.DefaultHtml}}')
    }
  )]
};
