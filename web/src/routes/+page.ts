import type { PageLoad } from './$types';

type GradeInfo = {
	student_id: string;
	grade: number | null;
};

type LatestData = {
	begin_at: number;
	end_at: number;
	semester: string;
	name: string;
	info: GradeInfo[];
};

export const load: PageLoad = async ({ fetch }) => {
	const res = await fetch(`/api/grade/latest`);
	const data: LatestData = await res.json();

	let submitted = 0;

	for (const info of data.info) {
		if (info.grade) {
			submitted += 1;
		}
	}

	return { ...data, submitted };
};
