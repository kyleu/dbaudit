package parse

import (
	"testing"
)

var x = []string{
	`exec sp_executesql N'select * from foo where x = @p0 and y = @p1;',N'@p0 decimal(29,10),@p1 nvarchar(4000)',@p0=1,@p1='bar'`,
	`exec sp_executesql N'select * from foo where x = @p0 and y = @p1;',N'@p0 uniqueidentifier,@p1 bit',@p0='00000000-0000-0000-0000-000000000000',@p1=1`,
	`exec sp_executesql N'
SELECT
	EventID,
	Title,
	DateTimeUtc,
	DoorOpenTimeUtc,
	EndTimeUtc,
	(SELECT COUNT(t.TicketID) FROM Outing ou JOIN Ticket t ON ou.OutingID = t.OutingID JOIN [Order] o ON t.OrderID = o.OrderID
		WHERE ou.EventID = ev.EventID AND o.DatePlacedUtc IS NOT NULL) AS NumberOfGuests,
	(SELECT COUNT(DISTINCT TicketID) FROM QRCodeScan WHERE EventID = ev.EventID) AS NumberOfCheckedInGuests,
	ISNULL(v.Timezone, org.TimezoneDefault) AS Timezone
FROM [Event] ev
JOIN Organization org ON ev.OrganizationID = org.OrganizationID
LEFT JOIN VenueMap vm ON ev.VenueMapID = vm.VenueMapID
LEFT JOIN Venue v ON vm.VenueID = v.VenueID
WHERE ev.OrganizationID IN (@orgIDs1) AND
ev.Deleted = 0 AND
(ev.DateTimeUtc BETWEEN DATEADD(DAY,-1, GETUTCDATE()) AND DATEADD(DAY,10, GETUTCDATE()) OR
	(ev.IsDateTimeTBA = 1 AND EXISTS(SELECT 1 FROM Ticket t JOIN Outing ou on t.OutingID = ou.OutingID
		           					 WHERE ou.EventID = ev.EventID AND
		           					t.ScanStartDateTime BETWEEN DATEADD(DAY,-1, GETUTCDATE()) AND DATEADD(DAY,10, GETUTCDATE()))))
ORDER BY DateTimeUtc',N'@orgIDs1 uniqueidentifier',@orgIDs1='C0169547-9495-4D36-8E8C-A6CF00F51D25'`,
}

func TestFoo(t *testing.T) {
	for _, test := range x {
		_, _, _, err := parseSQL(test)
		if err != nil {
			t.Error(err)
		}
	}
}
