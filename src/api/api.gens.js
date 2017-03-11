export const num = () => {
  const min = 0.0000;
  const max = 1.9000;
  return Math.random() * (max - min) + min;
}

export const token = (len) => {
  const possible = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789";
  let text = "";
  for (let i = 0; i < len; i++) {
    text += possible.charAt(Math.floor(Math.random() * possible.length));
  };
  return text;
}
