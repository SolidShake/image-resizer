import './style.css';
import './app.css';

import logo from './assets/images/gopher.png';
import {Greet} from '../wailsjs/go/main/App';

document.querySelector('#app').innerHTML = `
    <img id="logo" class="logo">
      <div class="result" id="result">Пожалуйста, выберите фотографии</div>
      <div class="input-box" id="input">
        <button class="btn" onclick="greet()">Выбрать</button>
      </div>
    </div>
`;
document.getElementById('logo').src = logo;

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
