export const parseStyles = (styles) => {
    return Object.entries(styles).reduce((acc, [key, value]) => {
      return `${acc}${key}: ${value};`;
    }, '');}