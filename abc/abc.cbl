       IDENTIFICATION DIVISION.
       PROGRAM-ID. ABC.

       DATA DIVISION.
       WORKING-STORAGE SECTION.
       01  WORKER.
         03  INP  PIC X(50).
         03  OUT  PIC X(50).
         03  TMPL PIC X(02).
         03  TMPN PIC X(01).
         03  LINE-1.
           05  A  PIC X(02).
           05  B  PIC X(02).
           05  C  PIC X(02).
         03  LINE-2.
           05  X  PIC X(01).
           05  Y  PIC X(01).
           05  Z  PIC X(01).

       PROCEDURE DIVISION.
           ACCEPT INP.
           UNSTRING INP
              DELIMITED BY SPACE
              INTO A, B, C 
           END-UNSTRING.

           ACCEPT INP.
           MOVE INP TO LINE-2.

           IF A > B
              MOVE A    TO TMPN
              MOVE B    TO A
              MOVE TMPN TO B
           END-IF.

           IF B > C
              MOVE B    TO TMPN
              MOVE C    TO B
              MOVE TMPN TO C
           END-IF.

           IF A > B
              MOVE A    TO TMPN
              MOVE B    TO A
              MOVE TMPN TO B
           END-IF.

           EVALUATE X ALSO Y ALSO Z
              WHEN "A" ALSO "B" ALSO "C"
                 CONTINUE
              WHEN "A" ALSO "C" ALSO "B"
                 MOVE C    TO TMPL 
                 MOVE B    TO C
                 MOVE TMPL TO B
              WHEN "B" ALSO "A" ALSO "C"
                 MOVE A    TO TMPL 
                 MOVE B    TO A
                 MOVE TMPL TO B
              WHEN "B" ALSO "C" ALSO "A"
                 MOVE C    TO TMPL 
                 MOVE A    TO C
                 MOVE B    TO A
                 MOVE TMPL TO B
              WHEN "C" ALSO "A" ALSO "B"
                 MOVE B    TO TMPL 
                 MOVE A    TO B
                 MOVE C    TO A
                 MOVE TMPL TO C
              WHEN "C" ALSO "B" ALSO "A"
                 MOVE A    TO TMPL 
                 MOVE C    TO A
                 MOVE TMPL TO C
           END-EVALUATE.

           STRING 
              A   DELIMITED BY SPACE
              " " DELIMITED BY SIZE 
              B   DELIMITED BY SPACE
              " " DELIMITED BY SIZE
              C   DELIMITED BY SPACE
              INTO OUT
           END-STRING.

           DISPLAY FUNCTION TRIM(OUT).
