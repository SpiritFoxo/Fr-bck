const express = require('express');
const path = require('path');

const app = express();
const PORT = process.env.PORT || 8080;

app.use(express.static(path.join(__dirname, '../public')));


app.get('/', (req, res) => {
    res.sendFile(path.join(__dirname, '../public/index.html'));
});

app.use((req, res) => {
    res.status(404).sendFile(path.join(__dirname, '../views/404.html'));
});

app.listen(PORT, () => {
    console.log(`Server is running on port ${PORT}`);
});
