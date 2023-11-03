import textwrap
import os
import subprocess


class NameGenerator:
    FEMALE_FILE = "https://www.cs.cmu.edu/afs/cs/project/ai-repository/ai/areas/nlp/corpora/names/female.txt"
    MALE_FILE = "https://www.cs.cmu.edu/afs/cs/project/ai-repository/ai/areas/nlp/corpora/names/male.txt"
    FAM_FILE = "https://www.cs.cmu.edu/afs/cs/project/ai-repository/ai/areas/nlp/corpora/names/other/family.txt"
    NAMES_FILE = "https://www.cs.cmu.edu/afs/cs/project/ai-repository/ai/areas/nlp/corpora/names/other/names.txt"
    TEMP_FAM_FILE = "./family.txt"
    TEMP_FEMALE_FILE = "./female.txt"
    TEMP_MALE_FILE = "./male.txt"
    TEMP_NAMES_FILE = "./names.txt"
    GO_FILE = "./names.go"

    def __init__(self):
        self.family_names = []
        self.female_names = []
        self.male_names = []
        self.all_names = []

    def download_files(self):
        """
        `download_files` downloads the files from the internet

        Args:
            None

        Returns:
            None
        """
        try:
            subprocess.run(
                ["wget", "--no-verbose",
                 self.FAM_FILE, self.FEMALE_FILE, self.MALE_FILE, self.NAMES_FILE]
            )
        except Exception as e:
            print("Error downloading files with error: ", e)

    def read_files(self):
        """
        `read_files` reads the files and returns the names

        Args:
            None

        Returns:
            None
        """
        with open(self.TEMP_FAM_FILE, "r+") as familyfile:
            family_names = [line.strip().replace(" ", "-")
                            for line in familyfile]
        self.family_names = list(set(family_names))

        with open(self.TEMP_FEMALE_FILE, "r+") as femalefile:
            female_names = [line.strip().replace(" ", "-")
                            for line in femalefile]
        self.female_names = list(set(female_names[6:]))

        with open(self.TEMP_MALE_FILE, "r+") as malefile:
            male_names = [line.strip().replace(" ", "-") for line in malefile]
        self.male_names = list(set(male_names[6:]))

        with open(self.TEMP_NAMES_FILE, "r+") as namesfile:
            all_names = [line.strip().replace(" ", "-") for line in namesfile]
        self.all_names = list(set(all_names))

    def print_summary(self):
        """
        `print_summary` prints the summary of the names

        Args:
            None

        Returns:
            None
        """
        print(f"""Number of female names is {len(self.female_names)}\n
        Number of male names is {len(self.male_names)}\n
        Number of family names is {len(self.family_names)}\n
        Number of general names is {len(self.all_names)}\n""")

    def generate_go_var(self, names, var_name):
        """
        `generate_go_var` generates the go variable

        Args:
            `names` (list): The list of names

            `var_name` (str): The name of the variable

        Returns:
            `list`: The list of lines of the go variable
        """

        var = f"{var_name} = []string{{\"" + \
            "\", \"".join(names) + "\"}"
        return textwrap.wrap(
            var, width=100, break_on_hyphens=False, break_long_words=False
        )

    def write_go_file(self):
        """
        `write_go_file` writes the go file

        Args:
            None

        Returns:
            None
        """
        wrapped_family_text = self.generate_go_var(
            self.family_names, "FamilyNames"
        )
        wrapped_female_text = self.generate_go_var(
            self.female_names, "FemaleNames"
        )
        wrapped_male_text = self.generate_go_var(
            self.male_names, "MaleNames"
        )
        wrapped_general_text = self.generate_go_var(
            self.all_names, "GeneralNames"
        )

        with open(self.GO_FILE, "w") as f:
            f.write("// Copyright (c) 0x6flab. All rights reserved.\n// \n")
            f.write("// SPDX-License-Identifier: GNU GENERAL PUBLIC LICENSE\n\n")

            f.write("// Code generated by namegenerator.py. DO NOT EDIT.\n\n")
            f.write("package namegenerator\n")
            f.write("var (\n")
            f.write("// FamilyNames is a list of family names\n")
            for line in wrapped_family_text:
                f.write(line + "\n")
            f.write("\n// FemaleNames is a list of female names\n")
            for line in wrapped_female_text:
                f.write(line + "\n")
            f.write("\n// MaleNames is a list of male names\n")
            for line in wrapped_male_text:
                f.write(line + "\n")
            f.write("\n// GeneralNames is a list of general names\n")
            for line in wrapped_general_text:
                f.write(line + "\n")
            f.write(")")

    def remove_file(self, file_path):
        """
        `remove_file` removes the temporary files

        Args:
            `file_path` (str): The path of the file to be removed

        Returns:
            None
        """
        os.remove(file_path)

    def run(self):
        """
        `run` runs the namegenerator. It downloads the files, reads the files,
        prints the summary, writes the go file and removes the temporary files.

        Args:
            None

        Returns:
            None
        """
        self.download_files()
        self.read_files()
        self.print_summary()
        self.write_go_file()
        temp_files = [self.TEMP_FAM_FILE, self.TEMP_FEMALE_FILE,
                      self.TEMP_MALE_FILE, self.TEMP_NAMES_FILE]
        for temp_file in temp_files:
            self.remove_file(temp_file)


def main():
    """
    Runs the namegenerator.
    """
    namegenerator = NameGenerator()
    namegenerator.run()


if __name__ == "__main__":
    main()
