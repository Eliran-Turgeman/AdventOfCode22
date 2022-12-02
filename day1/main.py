from __future__ import annotations


class ElfStorage:
    def __init__(self, items_calories: list[int]):
        self.calories = sum(items_calories)

    def __gt__(self, other):
        return self.calories > other.calories

    def __sub__(self, other):
        return self.calories - other.calories


class TopKElfs:
    def __init__(self, k: int):
        self.k = k
        self.top_elfs = [ElfStorage([]) for _ in range(k)]

    def add(self, elf: ElfStorage) -> None:
        index_to_add_elf = self._find_index_to_add_elf(elf)
        if index_to_add_elf is None:
            return

        self.top_elfs[index_to_add_elf] = elf

    def _find_index_to_add_elf(self, elf: ElfStorage) -> int | None:
        calories_diff = [elf - top_elf for top_elf in self.top_elfs]
        if all(diff <= 0 for diff in calories_diff):
            return None

        return calories_diff.index(max(calories_diff))

    def top_k_sum(self) -> int:
        return sum([elf.calories for elf in self.top_elfs])


if __name__ == '__main__':
    with open('input/input.txt', 'r') as puzzle_input:
        content = puzzle_input.readlines()
        content.append('\n')

    topK = TopKElfs(3)

    current_elf_calories: list[int] = []
    for line in content:
        if line == '\n':
            topK.add(ElfStorage(current_elf_calories))
            current_elf_calories = []
        else:
            current_elf_calories.append(int(line))

    print(f'Top Elfs Calories Sum: {topK.top_k_sum()}')
