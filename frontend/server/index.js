import express from "express";
import path from "path";
import { fileURLToPath } from "url";
import { dirname } from "path";
import process from "process";

const __filename = fileURLToPath(import.meta.url);
const __dirname = dirname(__filename);

const app = express();
const PORT = process.env.PORT || 3000;

app.use(express.static(path.join(__dirname, '..', 'dist')));

app.get('/', async (req, res) => {
    res.sendFile(path.join(__dirname, '..', 'dist', 'index.html'));
});
app.get('*', async (req, res) => {
    res.sendFile(path.join(__dirname, '..', 'dist', 'index.html'));
});

app.listen(PORT, () => {
  console.log(`Server is running on http://localhost:${PORT}`);
});
