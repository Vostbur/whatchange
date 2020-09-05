package main

const gladeTemplate = `<?xml version="1.0" encoding="UTF-8"?>
<!-- Generated with glade 3.36.0 -->
<interface>
  <requires lib="gtk+" version="3.22"/>
  <object class="GtkWindow" id="window_main">
    <property name="can_focus">False</property>
    <property name="title" translatable="yes">Folder watchdog</property>
    <property name="default_width">600</property>
    <property name="default_height">200</property>
    <child>
      <object class="GtkBox">
        <property name="visible">True</property>
        <property name="can_focus">False</property>
        <property name="orientation">vertical</property>
        <property name="baseline_position">bottom</property>
        <child>
          <object class="GtkScrolledWindow">
            <property name="visible">True</property>
            <property name="can_focus">True</property>
            <property name="shadow_type">in</property>
            <child>
              <object class="GtkTextView" id="textView">
                <property name="visible">True</property>
                <property name="can_focus">True</property>
                <property name="left_margin">5</property>
                <property name="right_margin">5</property>
                <property name="top_margin">5</property>
                <property name="bottom_margin">5</property>
              </object>
            </child>
          </object>
          <packing>
            <property name="expand">True</property>
            <property name="fill">True</property>
            <property name="position">0</property>
          </packing>
        </child>
        <child>
          <object class="GtkBox">
            <property name="visible">True</property>
            <property name="can_focus">False</property>
            <property name="margin_start">5</property>
            <property name="margin_end">5</property>
            <property name="margin_top">5</property>
            <property name="margin_bottom">5</property>
            <property name="spacing">5</property>
            <child>
              <object class="GtkButton" id="openButton">
                <property name="label" translatable="yes">Browse</property>
                <property name="visible">True</property>
                <property name="can_focus">True</property>
                <property name="receives_default">True</property>
              </object>
              <packing>
                <property name="expand">False</property>
                <property name="fill">True</property>
                <property name="position">0</property>
              </packing>
            </child>
            <child>
              <object class="GtkButton" id="stopButton">
                <property name="label" translatable="yes">Stop</property>
                <property name="visible">True</property>
                <property name="can_focus">True</property>
                <property name="receives_default">True</property>
              </object>
              <packing>
                <property name="expand">False</property>
                <property name="fill">True</property>
                <property name="pack_type">end</property>
                <property name="position">1</property>
              </packing>
            </child>
            <child>
              <object class="GtkButton" id="startButton">
                <property name="label" translatable="yes">Start</property>
                <property name="visible">True</property>
                <property name="can_focus">True</property>
                <property name="receives_default">True</property>
              </object>
              <packing>
                <property name="expand">False</property>
                <property name="fill">True</property>
                <property name="pack_type">end</property>
                <property name="position">2</property>
              </packing>
            </child>
          </object>
          <packing>
            <property name="expand">False</property>
            <property name="fill">True</property>
            <property name="position">1</property>
          </packing>
        </child>
      </object>
    </child>
    <child type="titlebar">
      <placeholder/>
    </child>
  </object>
</interface>
`
