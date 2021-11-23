const path = require("path");
const MiniCssExtractPlugin = require('mini-css-extract-plugin');


module.exports = {
    mode: "development",
    plugins: [new MiniCssExtractPlugin()],
    entry: './src/index.ts',
    module: {
        rules: [
            {
                test: /\.(s[ac]|c)ss$/i, // scss, sass and css
                use: [
                    MiniCssExtractPlugin.loader,
                    "css-loader",
                    "sass-loader",
                    "postcss-loader"
                ]
            },
            {
                test: /\.ts$/,
                exclude: /node_modules/,
                use: 'ts-loader',
                include: path.resolve(__dirname, 'src'),

            },
            {
                test: /\.js$/,
                exclude: /node_modules/,
                use: {
                    // without additional settings, this will reference .babelrc
                    loader: "babel-loader"
                }
            }
        ]
    },
    resolve: {
        extensions: ['.ts', '.js']
    },
    output: {
        filename: "bundle.js",
        path: path.resolve(__dirname, 'dist')
    },
    devServer: {
        static: {
            directory: path.join(__dirname, 'dist'),
        },
        compress: true,
        port: 8000,
    },
    devtool: 'eval-source-map',
}