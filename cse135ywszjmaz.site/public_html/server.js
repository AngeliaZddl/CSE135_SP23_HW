// app.js file

var jsonServer = require('json-server');

// Returns an Express server
var server = jsonServer.create();

// Set default middlewares (logger, static, cors and no-cache)
server.use(jsonServer.defaults());

// Add custom routes
// server.get('/custom', function (req, res) { res.json({ msg: 'hello' }) })

// Returns an Express router
var router = jsonServer.router('db.json');
// var staticRouter = jsonServer.router('static.json');
// var performanceRouter = jsonServer.router('performance.json');
// var activityRouter = jsonServer.router('activity.json');

server.use(router);
// server.use('/api/static', staticRouter);
// server.use('/api/performance', performanceRouter);
// server.use('/api/activity', activityRouter);

server.listen(3000);

var mysql = require('mysql');
var express = require('express');
var app = express();

var connection = mysql.createConnection({
  host: 'localhost',
  port: 3306,
  user: 'user',
  password: '',
  database: 'mydb'
})

app.get('/mydb', function(req, res) {  
    connection.query(queryString, function(err, rows, fields) {    
        if (err) throw err;    
        res.send(rows);  
    });
});