from flask import Flask, request, jsonify
from flask_cors import CORS

app = Flask(__name__)

#Habilitamos CORS
CORS(app)

#Para la base de datos utilizaremos un arreglo (en memoria)

users = []

@app.route('/users', methods=['GET'])
def get_users():
    return jsonify(users)

@app.route('/users', methods=['POST'])
def add_user():
    data = request.get_json()
    if 'name' and 'email' in data:
        users.append(data)
        return jsonify({'message': 'User added successfully'}), 201
    return jsonify({'error': 'Missing data'}), 400

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=5000)