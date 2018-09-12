import express from 'express';

function helloWorld(req, res) {
  res.status(200).send("<h1> Hello, world!</h1>");
}

const app = express();
app.get('/', helloWorld)

app.listen(3001);
