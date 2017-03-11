var gulp = require('gulp');
var cp = require('child_process');
var del = require('del');
var webpack = require('webpack-stream');
var webpackConfig = require('./webpack.config.js');
var nodemon = require('gulp-nodemon');
var rename = require('gulp-rename');
var livereload = require('gulp-livereload');
var exec = require('child_process').exec;

gulp.task('clean:build', function() {
  del('./public/js/*')
})

gulp.task('build', ['clean:build'], function() {
  return gulp.src('./src/main.js')
    .pipe(webpack(webpackConfig))
    .on('error', function handleError() {
      this.emit('end'); // Recover from errors
    })
    .pipe(gulp.dest('./'))
    .pipe(livereload());
});

gulp.task('watch:build', function() {
  livereload.listen();
  return gulp.watch('./src/**/*', ['build']);
});

// start backend
gulp.task('server:gin', function(done) {
  exec('./gin.server 8081 ./public', function (err, stdout, stderr) {
    console.log(stdout);
    console.log(stderr);
    done(err);
  });
});


// Main tasks
gulp.task('server', ['server:gin']);
gulp.task('watch', ['build', 'watch:build']);
gulp.task('default', ['watch','server']);
