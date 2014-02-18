var moment = require('moment');
var toRoman = require('roman-numerals').toRoman;
var util = require('util');

var now = moment();

var day = now.date();
// var currentLangData = moment.langData();
// var eng_day = currentLangData.ordinal(day);
var month = now.month() + 1;
var year = now.year();

var eng_date = moment().format("Do MMMM YYYY");
var us_date = moment().format("MMMM D YYYY");

var eng_date_roman = util.format("(%s.%s.%s)", toRoman(day), toRoman(month), toRoman(year));
var us_date_roman = util.format("(%s.%s.%s)", toRoman(month), toRoman(day), toRoman(year));

console.log("%s %s / %s %s", eng_date, eng_date_roman, us_date, us_date_roman);
