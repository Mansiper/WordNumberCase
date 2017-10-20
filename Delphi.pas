function WordNumberCase(Number: Integer; IfOne, IfTwo, IfFive: String; AddNumber: Boolean = False): String;
var
  Num, m: Integer;
begin
  Num := Number;
  If Number < 0 then
    Num := -Number;
  m := Num mod 10;
  Case m of
       1: Result:=IfOne;
    2..4: Result:=IfTwo;
  Else
    Result:=IfFive;
  End;
  If m in [1..4] then
    if Num mod 100 in [11..14] then
      Result:=IfFive;
	If AddNumber then
		Result := IntToStr(Number) + Result;
end;