// we can use objects as simple maps
const lookupTable = {
  'key': 'value',
  'differentTypes': 0
};

lookupTable['newVal'] = true;

console.log(lookupTable);
console.log(lookupTable['key']);

delete lookupTable.newVal;

console.log(lookupTable);

const obj = {
  data: 'test',
  hello: function() {
    console.log('data', this.data);
  }
};

// Note that we often use objects as maps, but
// JS provides a lot of flexibility - we couldn't
// use a golang map as an "object"
obj.hello();
