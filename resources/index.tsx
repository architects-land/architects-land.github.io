/* @refresh reload */
import { render } from "solid-js/web";
import "./scss/main.scss";
import App, { setupEvents } from "./App";

const root = document.getElementById("root");

render(() => <App />, root!);

setupEvents();
