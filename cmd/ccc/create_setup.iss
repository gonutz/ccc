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
ChangesEnvironment=true

[Languages]
Name: "english"; MessagesFile: "compiler:Default.isl"

[Tasks]
Name: "desktopicon"; Description: "{cm:CreateDesktopIcon}"; GroupDescription: "{cm:AdditionalIcons}"; Flags: unchecked
Name: "quicklaunchicon"; Description: "{cm:CreateQuickLaunchIcon}"; GroupDescription: "{cm:AdditionalIcons}"; Flags: unchecked; OnlyBelowVersion: 0,6.1
Name: envPath; Description: "Add to PATH variable" 

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

[Code]
{ Code taken from https://stackoverflow.com/questions/3304463/how-do-i-modify-the-path-environment-variable-when-running-an-inno-setup-install }
procedure EnvAddPath(instlPath: string);
var
    Paths: string;
begin
    { Retrieve current path (use empty string if entry not exists) }
    if not RegQueryStringValue(HKEY_CURRENT_USER, 'Environment', 'Path', Paths) then
        Paths := '';

    if Paths = '' then
        Paths := instlPath + ';'
    else
    begin
        { Skip if string already found in path }
        if Pos(';' + Uppercase(instlPath) + ';',  ';' + Uppercase(Paths) + ';') > 0 then exit;
        if Pos(';' + Uppercase(instlPath) + '\;', ';' + Uppercase(Paths) + ';') > 0 then exit;

        { Append App Install Path to the end of the path variable }
        Log(Format('Right(Paths, 1): [%s]', [Paths[length(Paths)]]));
        if Paths[length(Paths)] = ';' then
            Paths := Paths + instlPath + ';'  { don't double up ';' in env(PATH) }
        else
            Paths := Paths + ';' + instlPath + ';' ;
    end;

    { Overwrite (or create if missing) path environment variable }
    if RegWriteStringValue(HKEY_CURRENT_USER, 'Environment', 'Path', Paths)
    then Log(Format('The [%s] added to PATH: [%s]', [instlPath, Paths]))
    else Log(Format('Error while adding the [%s] to PATH: [%s]', [instlPath, Paths]));
end;

procedure EnvRemovePath(instlPath: string);
var
    Paths: string;
    P, Offset, DelimLen: Integer;
begin
    { Skip if registry entry not exists }
    if not RegQueryStringValue(HKEY_CURRENT_USER, 'Environment', 'Path', Paths) then
        exit;

    { Skip if string not found in path }
    DelimLen := 1;     { Length(';') }
    P := Pos(';' + Uppercase(instlPath) + ';', ';' + Uppercase(Paths) + ';');
    if P = 0 then
    begin
        { perhaps instlPath lives in Paths, but terminated by '\;' }
        DelimLen := 2; { Length('\;') }
        P := Pos(';' + Uppercase(instlPath) + '\;', ';' + Uppercase(Paths) + ';');
        if P = 0 then exit;
    end;

    { Decide where to start string subset in Delete() operation. }
    if P = 1 then
        Offset := 0
    else
        Offset := 1;
    { Update path variable }
    Delete(Paths, P - Offset, Length(instlPath) + DelimLen);

    { Overwrite path environment variable }
    if RegWriteStringValue(HKEY_CURRENT_USER, 'Environment', 'Path', Paths)
    then Log(Format('The [%s] removed from PATH: [%s]', [instlPath, Paths]))
    else Log(Format('Error while removing the [%s] from PATH: [%s]', [instlPath, Paths]));
end;

procedure CurStepChanged(CurStep: TSetupStep);
begin
    if (CurStep = ssPostInstall) and IsTaskSelected('envPath') then
     EnvAddPath(ExpandConstant('{app}'));
end;

procedure CurUninstallStepChanged(CurUninstallStep: TUninstallStep);
begin
    if CurUninstallStep = usPostUninstall then
      EnvRemovePath(ExpandConstant('{app}'));
end;
