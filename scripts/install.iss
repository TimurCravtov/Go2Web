[Setup]
AppName=Go2Web
AppVersion=1.0.0
DefaultDirName={autopf}\Go2Web
OutputDir=../Output
OutputBaseFilename=go2web
Compression=lzma
SolidCompression=yes
ArchitecturesInstallIn64BitMode=x64
; This line tells Windows to refresh environment variables after installation
ChangesEnvironment=yes

[Files]
Source: "go2web.exe"; DestDir: "{app}"; Flags: ignoreversion

[Icons]
Name: "{autoprograms}\Go2Web"; Filename: "{app}\go2web.exe"

[Registry]
; Adds the installation directory to the Current User's PATH variable safely
Root: HKCU; Subkey: "Environment"; ValueType: expandsz; ValueName: "Path"; ValueData: "{olddata};{app}"; Check: NeedsAddPath(ExpandConstant('{app}'))

[Code]
function NeedsAddPath(Param: string): boolean;
var
  OrigPath: string;
begin
  // Check if the Path registry key exists
  if not RegQueryStringValue(HKEY_CURRENT_USER, 'Environment', 'Path', OrigPath) then
  begin
    Result := True;
    exit;
  end;
  
  // Look for the path inside the existing Path variable. 
  // We add semicolons to ensure exact matches only.
  Result := Pos(';' + UpperCase(Param) + ';', ';' + UpperCase(OrigPath) + ';') = 0;
end;