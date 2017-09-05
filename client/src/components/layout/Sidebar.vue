<template>
  <aside class="menu app-sidebar animated" :class="{ slideInLeft: show, slideOutLeft: !show }">
    <p class="menu-label">
      General
    </p>
    <ul class="menu-list">
      <li>
        <a href="#">
          <span class="icon is-small is-angle">
            <i class="fa fa-angle-down"></i>
          </span>
        </a>
      </li>
    </ul>
  </aside>
</template>
