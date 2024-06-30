import Hero from "../../components/hero/Hero.tsx";

import data from "./seasons.json";
import List from "../../components/people/List.tsx";
import Card from "../../components/people/Card.tsx";
import { For } from "solid-js";

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
          <p></p>
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
