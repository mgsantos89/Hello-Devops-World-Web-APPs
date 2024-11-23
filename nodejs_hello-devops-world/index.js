// index.js
const express = require('express');
const os = require('os');
const app = express();
const port = 8080;

app.get('/', (req, res) => {
  const date = new Date();
  const message = `
    Hello Devops World!!!
    Date and Time: ${date.toLocaleString()}
    Hostname: ${os.hostname()}
  `;
  res.send(message);
});

app.listen(port, () => {
  console.log(`Server is running on http://localhost:${port}`);
});
