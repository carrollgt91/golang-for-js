import { publicFunc, privateFunc } from './mypackage/file.js';

// Runtime error
// /path/github.com/carrollgt91/golang-for-js/packages/index.js:5
// (0, _file.privateFunc)();
// TypeError: (0 , _file.privateFunc) is not a function
privateFunc();
publicFunc();
