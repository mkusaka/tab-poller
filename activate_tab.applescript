set _tab_was_found to false


      tell application "Safari"
                set _window_index to 1

      repeat with _window in windows
        try
          set _tab_count to (count of tabs in _window)
          set _tab_index to 1
          repeat with _tab in tabs of _window
                    if (url of _tab as string is "{.}") then
                set _tab_was_found to true
      activate
      tell _window
        set current tab to tab _tab_index
        set index to 1
      end tell
      exit repeat

      end if


            set _tab_index to _tab_index + 1
          end repeat
        end try
        set _window_index to _window_index + 1
      end repeat

      end tell

              if (_tab_was_found) then
        -- Bring window to front
        tell application "System Events" to tell process "Safari"
          perform action "AXRaise" of window 1
          -- account for instances when the window doesn't switch fast enough
          delay 0.5
          perform action "AXRaise" of window 1
          -- Prevent other running browsers from potentially activating
          return
        end tell
      end if





      tell application "Google Chrome"
                set _window_index to 1

      repeat with _window in windows
        try
          set _tab_count to (count of tabs in _window)
          set _tab_index to 1
          repeat with _tab in tabs of _window
                    if (url of _tab as string is "{{.}}") then
                set _tab_was_found to true
      activate
      tell _window
        set active tab index to  _tab_index
        set index to 1
      end tell
      exit repeat

      end if


            set _tab_index to _tab_index + 1
          end repeat
        end try
        set _window_index to _window_index + 1
      end repeat

      end tell

              if (_tab_was_found) then
        -- Bring window to front
        tell application "System Events" to tell process "Google Chrome"
          perform action "AXRaise" of window 1
          -- account for instances when the window doesn't switch fast enough
          delay 0.5
          perform action "AXRaise" of window 1
          -- Prevent other running browsers from potentially activating
          return
        end tell
      end if
