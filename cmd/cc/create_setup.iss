[Setup]
AppId={{8EF8C262-5FAC-4868-BD09-015574A964A4}
AppName=cc
AppVersion=1.0.0
AppPublisher=gonutz
AppPublisherURL=https://github.com/gonutz
AppSupportURL=https://github.com/gonutz
AppUpdatesURL=https://github.com/gonutz
DefaultDirName={pf}\cc
DefaultGroupName=cc
AllowNoIcons=yes
OutputDir=.
OutputBaseFilename="cc setup"
SetupIconFile=cc.ico
Compression=lzma
SolidCompression=yes
ChangesAssociations=yes

[Languages]
Name: "english"; MessagesFile: "compiler:Default.isl"

[Tasks]
Name: "desktopicon"; Description: "{cm:CreateDesktopIcon}"; GroupDescription: "{cm:AdditionalIcons}"; Flags: unchecked
Name: "quicklaunchicon"; Description: "{cm:CreateQuickLaunchIcon}"; GroupDescription: "{cm:AdditionalIcons}"; Flags: unchecked; OnlyBelowVersion: 0,6.1

[Files]
Source: "cc.exe"; DestDir: "{app}"; Flags: ignoreversion

[Icons]
Name: "{group}\cc"; Filename: "{app}\cc.exe"
Name: "{group}\{cm:UninstallProgram,cc}"; Filename: "{uninstallexe}"
Name: "{commondesktop}\cc"; Filename: "{app}\cc.exe"; Tasks: desktopicon
Name: "{userappdata}\Microsoft\Internet Explorer\Quick Launch\cc"; Filename: "{app}\cc.exe"; Tasks: quicklaunchicon

[Registry]
Root: HKCR; Subkey: ".cc"; ValueType: string; ValueName: ""; ValueData: "ccFile"; Flags: uninsdeletevalue
Root: HKCR; Subkey: "ccFile"; ValueType: string; ValueName: ""; ValueData: "cc xor file"; Flags: uninsdeletekey
Root: HKCR; Subkey: "ccFile\DefaultIcon"; ValueType: string; ValueName: ""; ValueData: "{app}\cc.exe,0"
Root: HKCR; Subkey: "ccFile\shell\open\command"; ValueType: string; ValueName: ""; ValueData: """{app}\cc.exe"" ""%1"""
