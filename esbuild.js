const esbuild = require('esbuild');
const postCssPlugin = require('esbuild-style-plugin');

const isWatchMode = process.env.MODE === 'watch';
const isProdMode = process.env.MODE !== 'dev' && !isWatchMode;

/** @type {esbuild.BuildContext} */
const bundleOptions = {
  entryPoints: ['frontend/app/app.tsx'],
  assetNames: 'assets/[name]-[hash]',
  outdir: 'public',
  bundle: true,
  minify: isProdMode,
  metafile: true,
  plugins: [
    postCssPlugin({
      postcss: {
        plugins: [require('tailwindcss'), require('autoprefixer')],
      },
    }),
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


