import Hero from "../../components/hero/Hero.tsx";

import data from "./seasons.json";
import List from "../../components/people/List.tsx";
import Card from "../../components/people/Card.tsx";
import { For, Show } from "solid-js";

import "./Season.scss";

export default function Season(props: any) {
  // @ts-ignore
  let info: any = data[props.id];
  return (
    <>
      <Hero
        title={info.name}
        image={info.image}
        description={info.description}
        min
      />
      <main>
        <div class="content--large large-text information">
          <h2>Information</h2>
          <For each={info.information.description}>
            {(item) => <p>{item}</p>}
          </For>
          <br />
          <p>
            Version de Minecraft : <strong>{info.information.version}</strong>
          </p>
          <p>
            Seed : <strong>{info.information.seed}</strong>
          </p>
          <p>
            Map :
            <strong>
              <Show when={info.information.map != "/"}>
                <a href={info.information.map} target={"_blank"}>
                  {info.information.map}
                </a>
              </Show>
              <Show when={info.information.map == "/"}>
                &nbsp;{info.information.map}
              </Show>
            </strong>
          </p>
          <p>
            Vidéo de présentation :{" "}
            <strong>
              <a href={info.information.presentationVideo} target={"_blank"}>
                {info.information.presentationVideo.slice(8)}
              </a>
            </strong>
          </p>
          <p class={"has-subtitle"}>Mods spécifiques à la saison :</p>
          <ul>
            <For each={info.information.mods}>
              {(item) => (
                <li>
                  <strong>{item}</strong>
                </li>
              )}
            </For>
          </ul>
        </div>
        <div class="content players">
          <h2>Liste des joueurs</h2>
          <List more>
            <For each={info.players}>
              {(item) => (
                <Card
                  name={item.name}
                  image={`/skins/${item.pseudo}.png`}
                  description={item.description}
                  link={item.link}
                />
              )}
            </For>
          </List>
        </div>
      </main>
    </>
  );
}
