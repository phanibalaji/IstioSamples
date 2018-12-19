'use strict';

const express = require('express');
const ip = require('ip');

// Constants
const PORT = 8080;
const HOST = '0.0.0.0';

// App
const app = express();
app.get('/upstreamNode', (req, res) => {
  res.setHeader('Content-Type', 'application/json');
  res.json({"Message": "Hello from Upstream Nodejs service",
            "version": "1",
            "Time": new Date(),
            "IP": ip.address()});
});

app.listen(PORT, HOST);
console.log(`Running on http://${HOST}:${PORT}`);
