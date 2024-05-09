const esbuild = require('esbuild');
const postCssPlugin = require('esbuild-style-plugin');
const copyPlugin = require('esbuild-plugin-copy').copy;

const isWatchMode = process.argv.includes('--watch');
const isProdMode = process.env.ENV === 'production';

/** @type {esbuild.BuildContext} */
const bundleOptions = {
  entryPoints: ['./app/app.tsx'],
  assetNames: 'assets/[name]-[hash]',
  outdir: 'public',
  bundle: true,
  minify: isProdMode,
  plugins: [
    postCssPlugin({
      postcss: {
        plugins: [require('tailwindcss'), require('autoprefixer')],
      },
    }),
    copyPlugin({
      assets: [
        {
          from: ['./assets/**'],
          to: ['./'],
        },
        {
          from: ['./favicons/**'],
          to: ['./'],
        },
      ],
    })
  ],
};

const onError = (error) => {
  console.error(`Build error: ${error}`)
  process.exit(1)
}

if (isWatchMode) {
  esbuild
    .context(bundleOptions)
    .then((ctx) => ctx.watch())
    .catch(onError)
} else {
  esbuild
    .build(bundleOptions)
    .catch(onError);
}
