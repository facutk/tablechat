import express from 'express';
const app = express();

app.get('/api/status', (req, res) => {
  res.json({ message: 'tablechat' });
});

const { HOST = '0.0.0.0', PORT = 3000 } = process.env;
app.listen(PORT, HOST, () => {
  console.log(`Server is running at http://${HOST}:${PORT}`);
});
