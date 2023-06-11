from bottle import Bottle, redirect, run

app = Bottle()

@app.route('/')
def forward():
    redirect('http://localhost:35901')

if __name__ == '__main__':
    run(app, host='localhost', port=8080)

