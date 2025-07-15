# Telegram Authentication Setup Guide

## Step 1: Create Telegram Bot

1. **Open Telegram** and search for `@BotFather`
2. **Send the command**: `/newbot`
3. **Follow the prompts**:
   - Choose a name for your bot (e.g., "DRPS Auth Bot")
   - Choose a username (e.g., "drpsauthbot")
4. **Save the bot token** that BotFather gives you (it looks like: `123456789:ABCdefGHIjklMNOpqrsTUVwxyz`)

## Step 2: Configure Environment Variables

1. **Create a `.env` file** in the `backend/` directory:
   ```bash
   cd backend
   cp env.template .env
   ```

2. **Edit the `.env` file** and add your Telegram bot token:
   ```env
   # JWT Configuration
   JWT_SECRET=your-super-secret-jwt-key-here
   
   # MongoDB Configuration
   MONGODB_URI=mongodb://localhost:27017
   
   # Telegram Bot Configuration
   TELEGRAM_SELLER_BOT_TOKEN=your-actual-telegram-bot-token-here
   TELEGRAM_CUSTOMER_BOT_TOKEN=your-actual-telegram-bot-token-here
   
   # Frontend URLs
   FRONTEND_URL=http://localhost:5174
   ```

## Step 3: Update Frontend Bot Username

1. **Open** `sellersDashboard/src/pages/AuthPage.vue`
2. **Find line 98** and update the bot username:
   ```javascript
   script.setAttribute('data-telegram-login', 'YOUR_BOT_USERNAME')
   ```
   Replace `YOUR_BOT_USERNAME` with your actual bot username (without the @ symbol)

## Step 4: Test the Setup

1. **Start the backend server**:
   ```bash
   cd backend
   go run main.go
   ```

2. **Start the frontend**:
   ```bash
   cd sellersDashboard
   npm run dev
   ```

3. **Navigate to** `http://localhost:5174/auth`
4. **Click the Telegram login button** and test the authentication

## Troubleshooting

### Common Issues:

1. **"Invalid Telegram signature" error**:
   - Check that your bot token is correct in the `.env` file
   - Ensure the bot username in the frontend matches your bot

2. **Telegram widget not loading**:
   - Verify your bot username is correct (without @ symbol)
   - Check browser console for JavaScript errors

3. **CORS errors**:
   - Ensure the frontend URL in `.env` matches your actual frontend URL
   - Check that the backend is running on port 8080

### Bot Configuration Tips:

1. **Set bot commands** (optional):
   - Send `/setcommands` to @BotFather
   - Add commands like: `start - Start the bot`

2. **Set bot description** (optional):
   - Send `/setdescription` to @BotFather
   - Add a description like: "Authentication bot for DRPS platform"

3. **Set bot about** (optional):
   - Send `/setabouttext` to @BotFather
   - Add about text like: "Official authentication bot for DRPS sellers"

## Security Notes:

- **Never commit your `.env` file** to version control
- **Keep your bot token secret** - it's like a password
- **Use different bots** for development and production
- **Regularly rotate your JWT_SECRET** in production

## Next Steps:

Once Telegram authentication is working, you can:
1. Set up Google OAuth (optional)
2. Customize the Telegram widget appearance
3. Add additional user profile fields
4. Implement email verification (optional) 