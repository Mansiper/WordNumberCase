function WordNumberCase(Number: Integer; IfOne, IfTwo, IfFive: String; AddNumber: Boolean = False): String;
begin
  Case Number mod 10 of
       1: Result := IfOne;
    2..4: Result := IfTwo;
  Else
    Result := IfFive;
  End;
  If Number MOD 100 in [11..14] then
    Result := IfFive;
	If AddNumber then
		Result := IntToStr(Number) + Result;
end;