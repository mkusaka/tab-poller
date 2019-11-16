set _output to ""


        tell application "Safari"
                  set _window_index to 1

        repeat with _window in windows
          try
            set _tab_count to (count of tabs in _window)
            set _tab_index to 1
            repeat with _tab in tabs of _window
              set _output to _output & "Safari:" &(_window_index as string) & ":" & (_tab_index as string) & "
" & url of _tab & "
" & name of _tab & "
"

              set _tab_index to _tab_index + 1
            end repeat
          end try
          set _window_index to _window_index + 1
        end repeat

        end tell






        tell application "Google Chrome"
                  set _window_index to 1

        repeat with _window in windows
          try
            set _tab_count to (count of tabs in _window)
            set _tab_index to 1
            repeat with _tab in tabs of _window
              set _output to _output & "Google Chrome:" &(_window_index as string) & ":" & (_tab_index as string) & "
" & url of _tab & "
" & title of _tab & "
"

              set _tab_index to _tab_index + 1
            end repeat
          end try
          set _window_index to _window_index + 1
        end repeat

        end tell




_output
