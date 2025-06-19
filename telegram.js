const TelegramBot = require('node-telegram-bot-api');
const axios = require('axios');
const TELEGRAM_TOKEN = '8149420340:AAHKAKmQOzMtgHUBW64Gb1awKxAHj6XJW-Y';

const bot = new TelegramBot(TELEGRAM_TOKEN, { polling: true });

bot.onText(/\/enviar (.+)/, async (msg, match) => {
  const chatId = msg.chat.id;
  const parts = match[1].split("|");
  if (parts.length !== 2) {
    return bot.sendMessage(chatId, "Formato: /enviar 5511999999999|Olá mundo!");
  }

  const number = parts[0].trim();
  const text = parts[1].trim();

  try {
    await axios.post("http://localhost:3000/send-text", {
      numbers: [number],
      message: text
    });
    bot.sendMessage(chatId, `✅ Mensagem enviada para ${number}`);
  } catch (err) {
    console.error(err);
    bot.sendMessage(chatId, `❌ Erro ao enviar: ${err.message}`);
  }
});

