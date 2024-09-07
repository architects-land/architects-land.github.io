import Hero from "../../components/hero/Hero.tsx";

import data from "../../assets/seasons.json";
import List from "../../components/people/List.tsx";
import Card from "../../components/people/Card.tsx";
import { For, Show } from "solid-js";

import "./Season.scss";

export type SeasonData = {
  name: string;
  description: string;
  image: string;
  logo: string;
  id: string;
  rp: boolean;
  players: PlayerData[];
  information: InformationData;
};

export type PlayerData = {
  name: string;
  description: string;
  pseudo: string;
  link: string;
  links: LinkData[];
  rp: string[];
};

export type LinkData = {
  link: string;
  name: string;
};

export type InformationData = {
  description: string[];
  version: string;
  mods: string[];
  video: string;
  videoThumbnail: string;
  seed: string;
  map: string;
};

export type Properties = {
  id: string;
};

export default function Season(props: any) {
  // @ts-ignore
  let info: SeasonData = data[props.id];
  console.log(data);
  document.title = `${info.name} - Architects Land`;
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
                  &nbsp;{info.information.map}
                </a>
              </Show>
              <Show when={info.information.map == "/"}>
                &nbsp;{info.information.map}
              </Show>
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
          <div class="information__presentation">
            <p>Vidéo de présentation :</p>
            <a href={info.information.video} target="_blank">
              <img
                src={info.information.videoThumbnail}
                alt="Miniature de la vidéo de présentation"
              />
            </a>
          </div>
        </div>
        <div class="content players">
          <h2>Liste des joueurs</h2>
          <List more>
            <Show when={!info.rp}>
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
            </Show>
            <Show when={info.rp}>
              <For each={info.players}>
                {(item) => (
                  <Card
                    name={item.name}
                    image={`/skins/${item.pseudo}.png`}
                    description={item.description}
                    link={`/season/${info.id}/player/${item.pseudo}`}
                  />
                )}
              </For>
            </Show>
          </List>
        </div>
      </main>
    </>
  );
}

export function getSeasons(): SeasonData {
  // @ts-ignore
  return data;
}
