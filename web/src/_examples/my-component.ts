import {html, LitElement} from 'lit';
import {customElement, property} from 'lit/decorators.js';

@customElement('my-component')
export class MyComponent extends LitElement {
  // Disable Shadow DOM and use Light DOM
  // This makes the element's content is part of the main DOM tree, accessible globally.
  // Styles from the parent document (global CSS) apply to the element's content.
  // Standard DOM APIs can access and manipulate the element's content.
  // Global event listeners can catch and handle the custom element triggered events.
  override createRenderRoot() {
    return this;
  }

  // @ts-ignore
  static styles = []

  @property()
  name = 'Somebody';

  city: string = '';

  override connectedCallback() {
    super.connectedCallback()
    this.getCity();
  }

  async getCity() {
    try {
      const response = await fetch('http://ip-api.com/json/');
      if (!response.ok) {
        throw new Error('Failed to fetch IP information');
      }

      const data = await response.json();
      this.city = data.city;
    } catch (error) {
      console.error('Error:', error);
    }
  }

  private handleClick(e: Event): void {
    // Because this is a Light DOM custom element, we need to stop the event propagation.
    e.preventDefault();
    alert(this.city);
  }

  override render() {
    return html`
      <ul role="list" class="grid gap-x-8 gap-y-12 sm:grid-cols-1 sm:gap-y-16 xl:col-span-2">
        <li>
          <div class="flex items-center gap-x-6">
            <img class="h-16 w-16 rounded-full"
                 src="https://images.unsplash.com/photo-1494790108377-be9c29b29330?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=facearea&facepad=2&w=256&h=256&q=80"
                 alt="">
            <div>
              <h3 class="text-base font-semibold leading-7 tracking-tight text-gray-900">${this.name}</h3>
            </div>
            <button
                class="rounded bg-white px-2 py-1 text-xs font-semibold text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 hover:bg-gray-50"
                @click="${this.handleClick}">
              Get City
            </button>
          </div>
        </li>
      </ul>
    `;
  }
}

declare global {
  interface HTMLElementTagNameMap {
    'my-component': MyComponent;
  }
}
