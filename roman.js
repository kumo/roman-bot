//
//  Roman - Twitter bot that tweets todays date in Roman Numerals format
//
// thanks to http://www.apcoder.com/2013/10/03/twitter-bot-20-minutes-node-js/ and http://www.apcoder.com/2013/10/15/targeted-twitter-bots-next-20-minutes/

var moment = require('moment');
var toRoman = require('roman-numerals').toRoman;
var util = require('util');

var Bot = require('./bot')
  , config1 = require('./config1');

var bot = new Bot(config1);

console.log('RomanBot: Running.');

Bot.prototype.today = function (callback) {
  var self = this;
 
  var now = moment();

  var day = now.date();
  var month = now.month() + 1;
  var year = now.year();

  var eng_date = moment().format("Do MMMM YYYY");
  var us_date = moment().format("MMMM D YYYY");

  var eng_date_roman = util.format("(%s.%s.%s)", toRoman(day), toRoman(month), toRoman(year));
  var us_date_roman = util.format("(%s.%s.%s)", toRoman(month), toRoman(day), toRoman(year));

  var roman_tweet = util.format("%s %s / %s %s", eng_date, eng_date_roman, us_date, us_date_roman);
  
  this.bot.tweet(roman_tweet, function (err, reply) {
    if(err) return callback(err);
  });
};

setInterval(function() {
  bot.today(function(err, reply)  {
    if(err) return handleError(err);

    console.log('\nTweet: ' + (reply ? reply.text : reply));
  });
}, 60 * 60 * 1000 * 24 );

function handleError(err) {
  console.error('response status:', err.statusCode);
  console.error('data:', err.data);
}

