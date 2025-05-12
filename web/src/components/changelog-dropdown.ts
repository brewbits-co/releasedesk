import {html, LitElement} from 'lit';
import {customElement, state} from 'lit/decorators.js';

@customElement('changelog-dropdown')
export class ChangelogDropdown extends LitElement {
  override createRenderRoot() {
    return this;
  }

  @state()
  isOpen = false;

  @state()
  addedEntries: Array<{type: string, label: string, bgColor: string, textColor: string, id: string}> = [];

  @state()
  changelogTypes = [
    { type: 'Changed', label: 'Changed', bgColor: 'bg-green-100', textColor: 'text-green-600' },
    { type: 'Added', label: 'Added', bgColor: 'bg-blue-100', textColor: 'text-blue-600' },
    { type: 'Fixed', label: 'Fixed', bgColor: 'bg-pink-100', textColor: 'text-pink-600' },
    { type: 'Removed', label: 'Removed', bgColor: 'bg-red-100', textColor: 'text-red-600' },
    { type: 'Security', label: 'Security', bgColor: 'bg-purple-100', textColor: 'text-purple-600' },
    { type: 'Deprecated', label: 'Deprecated', bgColor: 'bg-gray-200', textColor: 'text-gray-600' }
  ];

  toggleDropdown() {
    this.isOpen = !this.isOpen;
  }

  addChangelogEntry(type: string, label: string, bgColor: string, textColor: string) {
    // Generate a unique ID for this entry
    const id = `${type}_${Date.now()}`;

    // Add the new entry to the addedEntries array
    this.addedEntries = [...this.addedEntries, { type, label, bgColor, textColor, id }];

    // Close the dropdown
    this.isOpen = false;
  }

  override render() {
    return html`
      <div>
        <!-- Render all added entries -->
        ${this.addedEntries.map(entry => html`
          <div class="divide-y divide-gray-200 border-b border-gray-200">
            <div class="relative flex gap-3 py-2">
              <div class="w-1/12 flex justify-end">
                <span class="inline-flex items-center gap-x-1.5 rounded-md ${entry.bgColor} px-1.5 py-0.5 text-xs font-bold ${entry.textColor} h-6 mt-2">
                  ${entry.label}
                </span>
              </div>
              <div class="min-w-0 flex-1 text-sm/6">
                <input type="text"
                       name="${entry.type}"
                       id="${entry.id}"
                       autocomplete="off"
                       placeholder="Describe your change..."
                       class="block w-full bg-white px-3 py-1.5 text-base text-gray-900 placeholder:text-gray-400 outline-0 border-0 focus:ring-0 sm:text-sm/6">
              </div>
            </div>
          </div>
        `)}

        <!-- Add another change button and dropdown -->
        <div class="relative pt-2">
          <button type="button"
                  class="text-sm/6 font-semibold text-violet-800 hover:text-violet-700"
                  @click="${this.toggleDropdown}">
            <span aria-hidden="true">+</span> Add another change
          </button>

          ${this.isOpen ? html`
            <div class="absolute z-10 mt-2 w-56 origin-top-right rounded-md bg-white shadow-lg ring-1 ring-black ring-opacity-5 focus:outline-none">
              <div class="py-1">
                ${this.changelogTypes.map(type => html`
                  <button type="button"
                          class="block w-full px-4 py-2 text-left text-sm hover:bg-gray-100"
                          @click="${() => this.addChangelogEntry(type.type, type.label, type.bgColor, type.textColor)}">
                    <span class="inline-flex items-center gap-x-1.5 rounded-md ${type.bgColor} px-1.5 py-0.5 text-xs font-bold ${type.textColor}">
                      ${type.label}
                    </span>
                  </button>
                `)}
              </div>
            </div>
          ` : ''}
        </div>
      </div>
    `;
  }
}

declare global {
  interface HTMLElementTagNameMap {
    'changelog-dropdown': ChangelogDropdown;
  }
}
