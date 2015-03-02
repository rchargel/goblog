window.hljs.registerLanguage("yaml", function(hljs) {
  return {
    c: [
      {
        cN: 'string',
        b: '-.+$'
      },
      {
        cN: 'yamlattribute',
        b: ":.+$",
        v: [
          {
            cN: 'literal',
            bK: '\\d+(\\.\\d)?'
          },
          {
            cN: 'string',
            bK: '[^\\s]+$'
          }
        ]
      }
    ]
  };
});
