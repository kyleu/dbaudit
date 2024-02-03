// Content managed by Project Forge, see [projectforge.md] for details.
import "./client.css";
import {audit} from "./audit";
import {menuInit} from "./menu";
import {modeInit} from "./mode";
import {flashInit} from "./flash";
import {linkInit} from "./link";
import {timeInit} from "./time";
import {autocompleteInit} from "./autocomplete";
import {modalInit} from "./modal";
import {tagsInit} from "./tags";
import {formInit} from "./form";
import {themeInit} from "./theme";
import {appInit} from "./app";

declare global {
  interface Window { // eslint-disable-line @typescript-eslint/consistent-type-definitions
    "dbaudit": {
      relativeTime: (time: string, el?: HTMLElement) => string;
      autocomplete: (el: HTMLInputElement, url: string, field: string, title: (x: unknown) => string, val: (x: unknown) => string) => void;
      setSiblingToNull: (el: HTMLElement) => void;
      initForm: (frm: HTMLFormElement) => void;
      flash: (key: string, level: "success" | "error", msg: string) => void;
      tags: (el: HTMLElement) => void;
    };
    audit: (s: string, ...args: any) => void; // eslint-disable-line @typescript-eslint/no-explicit-any
  }
}

export function init(): void {
  const [s, i] = formInit();
  window.dbaudit = {
    relativeTime: timeInit(),
    autocomplete: autocompleteInit(),
    setSiblingToNull: s,
    initForm: i,
    flash: flashInit(),
    tags: tagsInit()
  };
  menuInit();
  modeInit();
  linkInit();
  modalInit();
  themeInit();
  window.audit = audit;
  appInit();
}

document.addEventListener("DOMContentLoaded", init);
