const { readFile } = require('node:fs/promises');
const { join } = require('node:path');

(async () => {
  const wasm = await WebAssembly.compile(
    await readFile(join(__dirname, 'module.wasm')),
  );
  const instance = await WebAssembly.instantiate(wasm, {
    myImports: {
      printFunc: param1 => {
        console.log(param1);
      }
    }
  });
})();