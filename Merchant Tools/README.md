# Accepting CloudCoins Automatically

Your servers can quickly receive authentic CloudCoins so that you can fill orders automatically. 

Your customers will send CloudCoins to your Skywallet using their CloudCoin Wallet or other system such as the SkyWallet ATM. 

After transfering money to your SkyWallet, your customers will need to send your servers a message specifying what they bought, how many CloudCoins they sent you, and what the "memo" was. 

Your servers will then "verify paymnet" to check your Skywallet to see if you received a transfer of that amount with the same memo.

Your server can call on our Payment Verifier exe to do all the work. Just tell the Payment Verifier, the amount, memo and a few other details. The Payment Verifier will return a true or false. 

The Payment Verifier is a executable program that runs on either Windows, Linux or Mac. It can be called from any langague that is able to call executables such as PHP, C#, JAVA, C++, Ruby, almost anything! 

The Payment Verifier is here in this Github folder. 

# Sending CloudCoins Automatically

You can send Cloudcoins quickly and reliably to your customer's Skywallet address by transferring them from your Skywallet account. 

You can call on our Transferrer executable to send the coins.  

The Transferrer is here in this Github folder. 

# Echoing the RAIDA

Before you verify payments or transfer funds, you may want to echo the RAIDA to make sure you hae connectivity. If you have a dependable Internet connection with no routers that block your ports, you will not need to echo. 

