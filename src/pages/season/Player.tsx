import { getSeasons, PlayerData } from "./Season.tsx";

import "./Season.scss";
import { For } from "solid-js";
import Button from "../../components/button/Button.tsx";

export type Properties = {
  id: string;
  player: string;
};

export default function Player(props: Properties) {
  // @ts-ignore
  const season = getSeasons()[props.id];
  // @ts-ignore
  let player: PlayerData = undefined;
  season.players.forEach((el: PlayerData) => {
    if (el.pseudo == props.player) player = el;
  });
  if (player == undefined) throw new Error("No valid player provided");
  return (
    <div class={"player-info"}>
      <div class={"player-info__skin"}>
        <img
          src={`/skins/${player.pseudo}.png`}
          alt={`Skin de ${player.pseudo}`}
        />
      </div>
      <div class={"player-info__description"}>
        <h3>{player.name}</h3>
        <p>{player.description}</p>
        <h4>RP</h4>
        <For each={player.rp}>{(item) => <p>{item}</p>}</For>
        <div class="player-info__links">
          <For each={player.links}>
            {(item) => <Button href={item.link} content={item.name} />}
          </For>
        </div>
      </div>
    </div>
  );
}
