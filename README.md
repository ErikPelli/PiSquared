# PiSquared
A Telegram bot that asks you a question and evaluate the response you provide.

Thanks to the labse_bert model, the evaluation of the answer is multilingual, so you can write the answer in your language and it will be checked correctly.

![Bot question example](/../photo/answer_example.png?raw=true "Bot question example")

## Download of model
In this repository there are some precompiled versions of the [hugging face importer tool provided by spaGO](https://github.com/nlpodyssey/spago/tree/main/cmd/huggingfaceimporter),
you can compile it using the instructions of original repository.
The total folder size will be around 5GB, but you can delete the pytorch_model.bin file to save 1.8GB with no problem.

### Download model using Windows
```
cd folderOfPiSquared
hf-importer-win64 --repo=models --model=pvl/labse_bert
```

### Download model using Linux
```
cd folderOfPiSquared
hf-importer-linux64 --repo=models --model=pvl/labse_bert
```

### Issues with Windows 11
I tested it using Windows and spago says that the OS doesn't support read only files.
I didn't find the solution so if you know how to resolve it, please open an issue.

## Configuration
Rename .env.example to .env

Fill the required fields:
- BOT_TOKEN is the Telegram token obtained by [BotFather](https://t.me/BotFather)
- QUESTIONS_FILE is the file that contains the questions and the answers for each subject
- SQLITE_DB is the SQLITE database file, and can be a path to a file or `file::memory:?cache=shared` for a temporary in memory database
- MODEL_FOLDER is the folder that contains the NPL model to check the answers

## Build this bot
Download Go compiler using a package manager (apt, snap, ecc) or from the [Go website](https://go.dev/dl/).
Then use this command in the project folder to compile without debug symbols and save space:
```
go build -ldflags "-s -w"
```
An executable file will be built and the settings are in the .env file.

## Thanks
- [SpaGO](https://github.com/nlpodyssey/spago), the ML/NPL library that processes the user answer
- [Gophercon SpaGO examples](https://github.com/matteo-grella/gophercon-eu-2021), that provides some useful examples to use SpaGO
- [Telebot](https://github.com/tucnak/telebot), a bot framework that this project use to connect with Telegram
- [Gorm](https://github.com/go-gorm/gorm), an ORM that abstracts SQL queries using Go data types