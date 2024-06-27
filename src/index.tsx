/* @refresh reload */
import { render } from 'solid-js/web'
import './scss/main.scss'
import App from './App'


const root = document.getElementById('root')

render(() => <App />, root!)

const seasons = document.getElementsByClassName('presentation__season')

for (const season of seasons) {
    season.addEventListener('mouseenter', _ => {
        season.classList.add("presentation__season--active")
    })
    season.addEventListener('mouseleave', _ => {
        season.classList.remove("presentation__season--active")
    })
}
