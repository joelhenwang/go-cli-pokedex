package utils

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"strings"

	"github.com/joelhenwang/go-cli-pokedex/structs"
)

func countLines(r io.Reader) (int, error) {
	buf := make([]byte, 32*1024)
	count := 0
	lineSep := []byte{'\n'}

	for {
		c, err := r.Read(buf)
		count += bytes.Count(buf[:c], lineSep)

		switch {
		case err == io.EOF:
			return count, nil

		case err != nil:
			return count, err
		}
	}
}

func LoadPokemonCsv(csvPath string) ([]structs.Pokemon, error) {
	f, err := os.Open(csvPath)
	if err != nil {
		return nil, err
	}

	defer f.Close()

	lines, err := countLines(f)
	if err != nil {
		return nil, err
	}

	_, err = f.Seek(0, io.SeekStart)
	if err != nil {
		return nil, err
	}

	f1, err := os.Open(csvPath)
	if err != nil {
		return nil, err
	}

	pokemon_list := make([]structs.Pokemon, 0, lines)

	scanner := bufio.NewScanner(f1)

	for scanner.Scan() {
		row := strings.Split(scanner.Text(), ",")

		new_pokemon := structs.Pokemon{
			Id:          row[1],
			Name:        row[2],
			Gen:         row[5],
			Type1:       row[9],
			Type2:       row[10],
			Hp:          row[18],
			Attack:      row[19],
			Defense:     row[20],
			SpAttack:    row[21],
			SpDefense:   row[22],
			Speed:       row[23],
			TotalPoints: row[17],
		}

		pokemon_list = append(pokemon_list, new_pokemon)
	}
	return pokemon_list, nil
}
