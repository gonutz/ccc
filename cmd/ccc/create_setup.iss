[Setup]
AppId={{FA5EB826-84C6-4F4E-ABF0-839F7B180F11}
AppName=ccc
AppVersion=1.0.0
AppPublisher=gonutz
AppPublisherURL=https://github.com/gonutz
AppSupportURL=https://github.com/gonutz
AppUpdatesURL=https://github.com/gonutz
DefaultDirName={pf}\ccc
DefaultGroupName=ccc
AllowNoIcons=yes
OutputDir=.
OutputBaseFilename="ccc setup"
SetupIconFile=ccc.ico
Compression=lzma
SolidCompression=yes
ChangesAssociations=yes
ArchitecturesInstallIn64BitMode=x64

[Languages]
Name: "english"; MessagesFile: "compiler:Default.isl"

[Tasks]
Name: "desktopicon"; Description: "{cm:CreateDesktopIcon}"; GroupDescription: "{cm:AdditionalIcons}"; Flags: unchecked
Name: "quicklaunchicon"; Description: "{cm:CreateQuickLaunchIcon}"; GroupDescription: "{cm:AdditionalIcons}"; Flags: unchecked; OnlyBelowVersion: 0,6.1

[Files]
Source: "ccc64.exe"; DestDir: "{app}"; DestName: "ccc.exe"; Check: Is64BitInstallMode; Flags: ignoreversion
Source: "ccc32.exe"; DestDir: "{app}"; DestName: "ccc.exe"; Check: not Is64BitInstallMode; Flags: ignoreversion

[Icons]
Name: "{group}\ccc"; Filename: "{app}\ccc.exe"
Name: "{group}\{cm:UninstallProgram,ccc}"; Filename: "{uninstallexe}"
Name: "{commondesktop}\ccc"; Filename: "{app}\ccc.exe"; Tasks: desktopicon
Name: "{userappdata}\Microsoft\Internet Explorer\Quick Launch\ccc"; Filename: "{app}\ccc.exe"; Tasks: quicklaunchicon

[Registry]
Root: HKCR; Subkey: ".ccc"; ValueType: string; ValueName: ""; ValueData: "cccFile"; Flags: uninsdeletevalue
Root: HKCR; Subkey: "cccFile"; ValueType: string; ValueName: ""; ValueData: "ccc file"; Flags: uninsdeletekey
Root: HKCR; Subkey: "cccFile\DefaultIcon"; ValueType: string; ValueName: ""; ValueData: "{app}\ccc.exe,0"
Root: HKCR; Subkey: "cccFile\shell\open\command"; ValueType: string; ValueName: ""; ValueData: """{app}\ccc.exe"" ""%1"""
Root: HKCR; Subkey: "*\shell\ccc this file\command"; ValueType: string; ValueName: ""; ValueData: """{app}\ccc.exe"" ""%1"""; Flags: uninsdeletekey