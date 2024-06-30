import "./People.scss";
import { Show } from "solid-js";

export default function List(props: any) {
  return (
    <>
      <Show when={props.more}>
        <div class={"content people__list large-text people__list--more"}>
          {props.children}
        </div>
      </Show>
      <Show when={!props.more}>
        <div class={"content people__list large-text"}>{props.children}</div>
      </Show>
    </>
  );
}
