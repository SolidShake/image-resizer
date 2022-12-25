import './style.css';
import './app.css';

import logo from './assets/images/gopher.png';
import {SavePath} from '../wailsjs/go/main/App';
import {Greet} from '../wailsjs/go/main/App';

document.querySelector('#app').innerHTML = `
    <img id="logo" class="logo">
      <div class="result" id="result">Пожалуйста, выберите фотографии</div>
      <div class="input-box" id="input">
        <button class="btn" onclick="savePath()">Выбрать папку сохранения</button>
        <button class="btn" onclick="greet()">Выбрать фотографии</button>
      </div>
    </div>
`;
document.getElementById('logo').src = logo;

// let nameElement = document.getElementById("name");
// nameElement.focus();
// let resultElement = document.getElementById("result");

// Setup the greet function
window.savePath = function () {
    // Get name
    // let name = nameElement.value;

    // Check if the input is empty
    // if (name === "") return;

    // Call App.Greet(name)
    try {
        // Greet(name)
        SavePath()
            .then((result) => {
                // Update result with data back from App.Greet()
                // resultElement.innerText = result;
            })
            .catch((err) => {
                console.error(err);
            });
    } catch (err) {
        console.error(err);
    }
};

// let nameElement = document.getElementById("name");
// nameElement.focus();
// let resultElement = document.getElementById("result");

// Setup the greet function
window.greet = function () {
    // Get name
    // let name = nameElement.value;

    // Check if the input is empty
    // if (name === "") return;

    // Call App.Greet(name)
    try {
        // Greet(name)
        Greet()
            .then((result) => {
                // Update result with data back from App.Greet()
                // resultElement.innerText = result;
            })
            .catch((err) => {
                console.error(err);
            });
    } catch (err) {
        console.error(err);
    }
};
