import {css, html, LitElement} from 'lit';
import {customElement} from 'lit/decorators.js';
import {getTailwindStyleSheet} from "../styles";

@customElement('my-input')
export class MyInput extends LitElement {
  // This element uses Shadow DOM to provide encapsulation.
  // The element's content is hidden inside a separate shadow DOM tree, not exposed to the global DOM.
  // External styles (global CSS) do not affect the shadow DOM content unless explicitly allowed.
  // The shadow DOM content is only accessible internally.
  // Shadow DOM block events from propagating up the DOM tree.

  // The getTailwindStyleSheet() imports the global CSS Stylesheet in order to Tailwind to be available.
  // @ts-ignore
  static styles = [getTailwindStyleSheet(), css`
      .internal-style {
          color: brown;
      }
  `]

  override render() {
    return html`
      <label for="email" class="block text-sm font-medium leading-6 text-gray-900 internal-style">
        <slot></slot>
      </label>
      <div class="mt-2">
        <input type="email" name="email" id="email"
               class="block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-900 sm:text-sm sm:leading-6"
        >
      </div>
    `;
  }
}

declare global {
  interface HTMLElementTagNameMap {
    'my-input': MyInput;
  }
}
