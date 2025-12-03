<script lang="ts">
  let {
    showInnerCircle,
    onclose,
  }: { showInnerCircle: boolean; onclose: () => void } = $props();

  let activeTab = $state<'party' | 'innerCircle' | 'about'>('party');

  const YEARS = [
    '2016',
    '2017',
    '2018',
    '2019',
    '2020',
    '2021',
    '2022',
    '2023',
    '2024',
    '2025',
  ];

  const tabs: { [key in typeof activeTab]: string } = $derived({
    party: 'Hochzeits&shy;feier',
    innerCircle: showInnerCircle ? 'Standes&shy;amt' : '',
    about: 'Über uns',
  });

  function setActiveTab(tab: string) {
    if (tab === 'party' || tab === 'innerCircle' || tab === 'about') {
      if (tab === 'innerCircle' && !showInnerCircle) {
        return;
      }
      activeTab = tab;
    } else {
      activeTab = 'party';
    }
  }
</script>

<header class="flex items-center justify-between border-b-2 border-blue-500">
  <nav class="flex">
    {#each Object.entries(tabs) as [tab, tabLabel]}
      {#if tabLabel}
        <button
          class={[
            'cursor-pointer px-4 py-2 transition-colors lg:px-8 lg:py-4',
            activeTab === tab
              ? 'bg-blue-500 text-white hover:bg-blue-500/90'
              : 'text-blue-500 hover:bg-blue-500/10',
          ]}
          onclick={() => setActiveTab(tab)}>{@html tabLabel}</button
        >
      {/if}
    {/each}
  </nav>
  <button
    class="cursor-pointer px-4 py-2 text-blue-500 transition-colors hover:bg-blue-500/10 lg:px-8 lg:py-4"
    onclick={() => onclose()}>x</button
  >
</header>
<div class="flex-1 overflow-y-auto p-4">
  {#if activeTab === 'party'}
    <div
      class="grid grid-cols-1 gap-8 font-normal text-blue-500 lg:grid-cols-3"
    >
      <div class="flex flex-col gap-4">
        <div>
          <h3 class="text-xl font-bold text-pink-500">Wo?</h3>
          <address class="not-italic">
            <a href="https://foodlab.hamburg/" target="_blank" class="underline"
              ><h4 class="font-bold">foodlab</h4></a
            >
            Überseeallee 10<br />
            20457 Hamburg
          </address>
        </div>
        <div>
          <h3 class="text-xl font-bold text-pink-500">Anfahrt</h3>
          <ul class="flex flex-col gap-2">
            <li>
              <h4 class="font-bold">Zu Fuß</h4>
              <p>
                Ihr könnt in etwa 25 Minuten vom Hauptbahnhof bis zum foodlab
                laufen.
              </p>
            </li>
            <li>
              <h4 class="font-bold">ÖPNV</h4>
              <p>
                Die U4 Richtung Elbbrücken bis Überseequartier nehmen. Vom
                Hauptbahnhof dauert das gerade einmal 5 Minuten. Dann müsst ihr
                nur noch ein paar Meter laufen.
              </p>
            </li>
            <li>
              <h4 class="font-bold">Auto</h4>
              <p>
                Da das Parken in der Hafencity ziemlich schwierig (und teuer)
                ist, empfehlen wir euch in jedem Fall den ÖPNV zu nutzen.
              </p>
              <p>
                Falls ihr doch mit dem Auto kommt, könnt ihr in der Nähe z.B. im
                Parkhaus Elbarkaden oder im Westfield Parkhaus unterkommen.
              </p>
            </li>
          </ul>
        </div>
        <div>
          <h3 class="text-xl font-bold text-pink-500">Hotelvorschläge</h3>
          <ul class="flex flex-col gap-2">
            <li>
              <a
                href="https://25hours-hotels.com/hamburg/hafencity/"
                target="_blank"
                class="underline"><h4 class="font-bold">25hours</h4></a
              >
              €€€ | ~400m
            </li>
            <li>
              <a
                href="https://www.ihg.com/holidayinn/hotels/gb/en/hamburg/hamhc/hoteldetail"
                target="_blank"
                class="underline"
                ><h4 class="font-bold">Holiday Inn Speicherstadt</h4></a
              >
              €€ | ~450m
            </li>
            <li>
              <a
                href="https://www.premierinn.com/de/de/hotels/deutschland/hamburg/hamburg/hamburg-city-zentrum.html"
                target="_blank"
                class="underline"
                ><h4 class="font-bold">
                  Premier Inn Hamburg City (Zentrum)
                </h4></a
              >
              €€ | ~1,1km
            </li>
          </ul>
        </div>
      </div>
      <div class="flex flex-col gap-4">
        <div>
          <h3 class="text-xl font-bold text-pink-500">Wann?</h3>
          <p>22.05.2026</p>
          <p>15 Uhr</p>
          <p class="italic">
            Plant am besten bis 14:30 Uhr da zu sein, sodass wir die Trauung gut
            um 15 Uhr beginnen können.
          </p>
        </div>
        <div>
          <h3 class="text-xl font-bold text-pink-500">Ablauf</h3>
          <ul class="flex flex-col gap-4">
            <li>
              <strong>15:00 Uhr</strong>
              <p>Freie Trauung 💍</p>
            </li>
            <li>
              <strong>15:45 Uhr</strong>
              <p>Hugs 💓</p>
            </li>
            <li>
              <strong>16:00 Uhr</strong>
              <p>Empfang 🥂</p>
            </li>
            <li>
              <strong>18:30 Uhr</strong>
              <p>Essen 🍜</p>
            </li>
            <li>
              <strong>20:30 Uhr</strong>
              <p>Party 🎉</p>
            </li>
            <span class="text-sm">Der Ablauf ist noch nicht final.</span>
          </ul>
        </div>
      </div>
      <div class="flex flex-col gap-4">
        <div>
          <h3 class="text-xl font-bold text-pink-500">Geschenke</h3>
          <p>
            Wir würden uns sehr darüber freuen, wenn ihr einen Teil zu unseren
            Flitterwochen beitragen würdet.
          </p>
          <p>Wir planen eine Reise nach <strong>New York</strong>!</p>
          <p>
            Wenn ihr mögt (auch wenn es nicht so charmant ist), wäre es am
            besten, wenn ihr das Geld direkt über PayPal oder per Überweisung
            schickt.
          </p>
          <div class="mt-2">
            <a href="https://paypal.com" target="_blank" class="underline"
              >Zu PayPal</a
            >
            <p>IBAN: DE50 5001 0517 5429 5063 95</p>
          </div>
        </div>
        <div>
          <h3 class="text-xl font-bold text-pink-500">Dresscode</h3>
          <p class="italic">Summery Cocktail Attire</p>
          <p>
            Darunter verstehen wir: Anzug mit Krawatte oder Kleider mit
            Midi-Länge in sommerlichen Farben / Prints. Lasst euch doch von den
            Farben der Einladung inspirieren!
          </p>
        </div>
        <div>
          <h3 class="text-xl font-bold text-pink-500">Reden / Beiträge</h3>
          <p>
            Bitte meldet euch bei den Trauzeug*innen, wenn ihr eine Rede oder
            Ähnliches zur Feier beitragen möchtet. Das hilft uns bei der
            Planung.
          </p>
          <ul class="mt-2">
            <li>
              Luisa Woitera <a class="italic" href="tel:+4915168806999"
                >+49 151 68806999</a
              >
            </li>
            <li>
              Nicolas Barfurth <a class="italic" href="tel:+4915150246804"
                >+49 151 50246804</a
              >
            </li>
          </ul>
        </div>
      </div>
    </div>
  {:else if activeTab === 'innerCircle'}
    <div
      class="grid grid-cols-1 gap-4 font-normal text-blue-500 lg:grid-cols-3"
    >
      <div class="flex flex-col gap-4">
        <div>
          <h3 class="text-xl font-bold text-pink-500">Wo?</h3>
          <address class="not-italic">
            <a
              href="https://maps.app.goo.gl/JrkBH6pmdyA3LLXT7"
              target="_blank"
              class="underline"
              ><h4 class="font-bold">Standesamt Hamburg-Nord</h4></a
            >
            Robert-Koch-Straße 17<br />
            20249 Hamburg-Eppendorf
          </address>
        </div>
      </div>
      <div class="flex flex-col gap-4">
        <div>
          <h3 class="text-xl font-bold text-pink-500">Wann?</h3>
          <p>15.05.2026</p>
          <p>11:00 Uhr</p>
        </div>
      </div>
      <div class="flex flex-col gap-4">
        <div>
          <h3 class="text-xl font-bold text-pink-500">Ablauf</h3>
          <p>Die Trauung dauert ca. 20 Minuten.</p>
          <p>Anschließend gehen wir als Gruppe gemeinsam zum Mittagessen.</p>
          <p class="text-sm italic">Das Restaurant steht noch nicht fest.</p>
        </div>
      </div>
    </div>
  {:else if activeTab === 'about'}
    <div class="flex flex-col gap-4 text-blue-500">
      <h3 class="text-xl font-bold text-pink-500">10 Jahre später...</h3>
      <p>Obwohl ihr uns ja alle kennt hier nochmal ein Mini-Rückblick:</p>
      <p>
        Wir haben uns 2016 während des Abiturs in Stade am Athenaeum
        kennengelernt. Wir sind also quasi so eine "Schulromanze" wie aus
        kitschigen Liebesfilmen - aber halt in echt. Nach dem Abitur ging
        Sebastian für sein Studium nach Stuttgart und Sarah machte ein Jahr als
        Au-Pair in London. Ein Jahr Fernbeziehung haben wir aber auch
        überstanden. Danach kam Sarah ebenfalls nach Stuttgart zum Studieren.
        Seitdem wohnen wir gemeinsam - zunächst in Stuttgart (und Umgebung), und
        seit 2022 in München.
      </p>
      <p>
        Nach fast genau 10 Jahren wollen wir jetzt dann (endlich) heiraten.
        Damit die passende Stimmung aufkommt, hier noch ein paar Eindrücke aus
        unseren gemeinsamen Jahren, die hierher geführt haben.
      </p>
      <p>
        Wir freuen uns auf euch,<br />
        <span class="text-pink-500">Sarah & Sebastian</span>
      </p>
      <div class="grid gap-4 lg:grid-cols-2">
        {#each YEARS as year}
          <div>
            <h3 class="text-xl font-bold text-pink-500">
              {year}
            </h3>
            <img
              src="{import.meta.env
                .VITE_YEARS_IMAGES_BASE_URL}{year}.{import.meta.env
                .VITE_YEARS_IMAGES_EXT}"
              alt="Bild aus dem Jahr {year}"
              class="w-full rounded-lg border-2 border-blue-500"
            />
          </div>
        {/each}
      </div>
    </div>
  {/if}
</div>
